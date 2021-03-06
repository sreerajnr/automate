#!/usr/bin/env bash

install_chefdk() {
    local channel="${1:-stable}"
    local version="${2:-latest}"
    curl https://omnitruck.chef.io/install.sh | bash -s -- -c "$channel" -P chefdk -v "$version"
    chef --version
}

start_requestbin() {
    export GOBIN="/go/bin"
    go install -mod=vendor integration/helpers/requestbin/requestbin.go
    cat > /etc/systemd/system/requestbin.service <<EOF
[Unit]
Description=Requestbin

[Service]
ExecStart=/go/bin/requestbin

[Install]
WantedBy=default.target
EOF
    systemctl start requestbin
}
start_loadbalancer() {
    export GOBIN="/go/bin"
    go install -mod=vendor ./integration/helpers/loadbalancer/
    #shellcheck disable=SC2154
    cat > /etc/systemd/system/loadbalancer.service <<EOF
[Unit]
Description=Load Balancer

[Service]
ExecStart=/go/bin/loadbalancer $(pwd)/$test_config_path $@

[Install]
WantedBy=default.target
EOF
    systemctl start loadbalancer.service
}

# Start elasticsearch, if an S3 endpoint argument is passed in then enable it
start_external_elasticsearch() {
  local s3_endpoint
  local version
  version="6.8.3"
  s3_endpoint=$1

  if [ ! -f elasticsearch.rpm ]; then
    curl -o elasticsearch.rpm "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-${version}.rpm"
  fi
  rpm --install elasticsearch.rpm

  adduser hab

  mkdir -p /var/opt/chef-automate/backups/automate-elasticsearch-data
  chmod -R 0777 /var/opt/chef-automate/backups

  mkdir -p /var/run/elasticsearch
  chown elasticsearch:elasticsearch /var/run/elasticsearch
  chown -R elasticsearch:elasticsearch /usr/share/elasticsearch/

  cat > /etc/elasticsearch/elasticsearch.yml <<EOF
cluster.name: "external-network"
network.host: 127.0.0.1
http.port: 59200
transport.tcp.port: "59300-59400"
path.repo: "/var/opt/chef-automate/backups"
discovery.zen.minimum_master_nodes: 1
EOF

  if [[ -n "${s3_endpoint}" ]]; then
    echo "y" | /usr/share/elasticsearch/bin/elasticsearch-plugin install "https://artifacts.elastic.co/downloads/elasticsearch-plugins/repository-s3/repository-s3-${version}.zip"

    cat >> /etc/elasticsearch/elasticsearch.yml <<EOF
s3.client.default.protocol: "https"
s3.client.default.read_timeout: "50s"
s3.client.default.max_retries: 3
s3.client.default.use_throttle_retries: true
s3.client.default.endpoint: "${s3_endpoint}"
EOF

    echo "$AWS_ACCESS_KEY_ID" | /usr/share/elasticsearch/bin/elasticsearch-keystore add s3.client.default.access_key
    echo "$AWS_SECRET_ACCESS_KEY" | /usr/share/elasticsearch/bin/elasticsearch-keystore add s3.client.default.secret_key
    echo "$AWS_SESSION_TOKEN" | /usr/share/elasticsearch/bin/elasticsearch-keystore add s3.client.default.session_token
  fi

  systemctl start elasticsearch.service
}
