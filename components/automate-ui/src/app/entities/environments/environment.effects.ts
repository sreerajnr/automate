import { Injectable } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { Actions, Effect, ofType } from '@ngrx/effects';
import { of as observableOf } from 'rxjs';
import { catchError, mergeMap, map } from 'rxjs/operators';
import { CreateNotification } from 'app/entities/notifications/notification.actions';
import { Type } from 'app/entities/notifications/notification.model';

import {
  GetEnvironments,
  GetEnvironmentsSuccess,
  EnvironmentsSuccessPayload,
  GetEnvironmentsFailure,
  EnvironmentActionTypes
} from './environment.action';

import { EnvironmentRequests } from './environment.requests';

@Injectable()
export class EnvironmentEffects {
  constructor(
    private actions$: Actions,
    private requests: EnvironmentRequests
  ) { }

  @Effect()
  getEnvironments$ = this.actions$.pipe(
      ofType(EnvironmentActionTypes.GET_ALL),
      mergeMap(({ payload: { server_id, org_id } }: GetEnvironments) =>
        this.requests.getEnvironments(server_id, org_id).pipe(
          map((resp: EnvironmentsSuccessPayload) => new GetEnvironmentsSuccess(resp)),
          catchError((error: HttpErrorResponse) =>
          observableOf(new GetEnvironmentsFailure(error))))));

  @Effect()
  getEnvironmentsFailure$ = this.actions$.pipe(
      ofType(EnvironmentActionTypes.GET_ALL_FAILURE),
      map(({ payload }: GetEnvironmentsFailure) => {
        const msg = payload.error.error;
        return new CreateNotification({
          type: Type.error,
          message: `Could not get environments: ${msg || payload.error}`
        });
      }));

}
