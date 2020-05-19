import { ComponentFixture, TestBed } from '@angular/core/testing';
import { UserTableComponent } from './user-table.component';
import { RouterTestingModule } from '@angular/router/testing';
import { MockComponent } from 'ng2-mock-component';

const baseUrl = '/some/path';

describe('UserTableComponent', () => {
  let component: UserTableComponent;
  let fixture: ComponentFixture<UserTableComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [
        RouterTestingModule
      ],
      declarations: [
        MockComponent({ selector: 'app-authorized', inputs: ['allOf'] }),
        MockComponent({ selector: 'chef-toolbar' }),
        UserTableComponent
      ]
    })
    .compileComponents();
    fixture = TestBed.createComponent(UserTableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should be created', () => {
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });

  describe('when {get,create}PermissionsPath are not passed', () => {
    it('should populate the permission paths using baseUrl', () => {
      component.baseUrl = baseUrl;
      component.ngOnInit();
      fixture.detectChanges();
      expect(component.getPermissionsPath).toEqual([baseUrl, 'get']);     // TODO fails
      expect(component.createPermissionsPath).toEqual([baseUrl, 'post']); // TODO fails
    });
  });

  describe('when {get,create}PermissionsPath are not passed', () => {
    it('should not populate the permission paths using baseUrl', () => {
      const getPath = ['/some/parameterized/{path}', 'get', 'my-path'];
      const postPath = ['/some/parameterized/{path}', 'post', 'my-path'];
      component.getPermissionsPath = getPath;
      component.createPermissionsPath = postPath;
      component.ngOnInit();
      fixture.detectChanges();
      expect(component.getPermissionsPath).toEqual(getPath);
      expect(component.createPermissionsPath).toEqual(postPath);
    });
  });
});
