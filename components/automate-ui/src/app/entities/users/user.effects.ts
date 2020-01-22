import { catchError, mergeMap, map, filter } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { HttpErrorResponse } from '@angular/common/http';
import { Store } from '@ngrx/store';
import { Actions, Effect, ofType } from '@ngrx/effects';
import { of, combineLatest } from 'rxjs';
import { identity } from 'lodash/fp';

import { HttpStatus } from 'app/types/types';
import { NgrxStateAtom } from 'app/ngrx.reducers';
import { CreateNotification } from 'app/entities/notifications/notification.actions';
import { Type } from 'app/entities/notifications/notification.model';
import { iamMajorVersion } from 'app/entities/policies/policy.selectors';
import { IAMMajorVersion } from 'app/entities/policies/policy.model';
import {
  CreateUser,
  CreateUserSuccess,
  CreateUserFailure,
  DeleteUser,
  DeleteUserSuccess,
  DeleteUserFailure,
  GetUsers,
  GetUsersSuccess,
  GetUsersSuccessPayload,
  GetUsersFailure,
  GetUser,
  GetUserSuccess,
  GetUserFailure,
  UpdateUser,
  UpdateUserSuccess,
  UpdateUserFailure,
  UpdateSelf,
  UpdateSelfSuccess,
  UserActionTypes
} from './user.actions';
import { UserRequests } from './user.requests';
import { User } from './user.model';
import { SelfUser } from './userself.model';

@Injectable()
export class UserEffects {
  constructor(
    private actions$: Actions,
    private requests: UserRequests,
    private store$: Store<NgrxStateAtom>
  ) { }

  @Effect()
  getUsers$ = combineLatest([
    this.actions$.pipe(ofType<GetUsers>(UserActionTypes.GET_ALL)),
    this.store$.select(iamMajorVersion).pipe(filter(identity))])
    .pipe(
      mergeMap(([_action, version]: [GetUsers, IAMMajorVersion]) =>
        this.requests.getUsers(version).pipe(
          map((resp: GetUsersSuccessPayload) => new GetUsersSuccess(resp)),
          catchError((error: HttpErrorResponse) => of(new GetUsersFailure(error))))));

  @Effect()
  getUsersFailure$ = this.actions$.pipe(
    ofType(UserActionTypes.GET_ALL_FAILURE),
    map(({ payload }: GetUsersFailure) => {
        const msg = payload.error.error;
        return new CreateNotification({
        type: Type.error,
        message: `Could not get users: ${msg || payload.error}`
      });
    }));

  @Effect()
  getUser$ = combineLatest([
    this.actions$.pipe(ofType<GetUser>(UserActionTypes.GET)),
    this.store$.select(iamMajorVersion).pipe(filter(identity))])
    .pipe(
      mergeMap(([action, version]: [GetUser, IAMMajorVersion]) =>
        this.requests.getUser(action.payload.id, version).pipe(
          map((resp: User) => new GetUserSuccess(resp)),
          catchError((error: HttpErrorResponse) => of(new GetUserFailure(error))))));

  @Effect()
  getUserFailure$ = this.actions$.pipe(
    ofType(UserActionTypes.GET_FAILURE),
    map(({ payload: { error } }: GetUserFailure) => {
      const msg = error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not get user: ${msg || error}`
      });
    }));

  @Effect()
  updateUser$ = combineLatest([
    this.actions$.pipe(ofType<UpdateUser>(UserActionTypes.UPDATE)),
    this.store$.select(iamMajorVersion).pipe(filter(identity))])
    .pipe(
      mergeMap(([action, version]: [UpdateUser, IAMMajorVersion]) =>
      this.requests.updateUser(action.payload, version).pipe(
        map((resp: User) => new UpdateUserSuccess(resp)),
        catchError((error: HttpErrorResponse) => of(new UpdateUserFailure(error))))));

  @Effect()
  updateUserSuccess$ = this.actions$.pipe(
    ofType(UserActionTypes.UPDATE_SUCCESS),
    map(() => new CreateNotification({
      type: Type.info,
      message: 'Successfully updated user'
    })));

  @Effect()
  updateUserFailure$ = this.actions$.pipe(
    ofType(UserActionTypes.UPDATE_FAILURE),
    map(({ payload: { error } }: UpdateUserFailure) => {
      const msg = error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not update user: ${msg || error}`
      });
    }));

  @Effect()
  updateSelf$ = combineLatest([
    this.actions$.pipe(ofType<UpdateSelf>(UserActionTypes.UPDATE_SELF)),
    this.store$.select(iamMajorVersion).pipe(filter(identity))])
    .pipe(
      mergeMap(([action, version]: [UpdateSelf, IAMMajorVersion]) =>
      this.requests.updateSelf(action.payload, version).pipe(
        map((resp: SelfUser) => new UpdateSelfSuccess(resp)),
        catchError((error: HttpErrorResponse) => of(new UpdateUserFailure(error))))));

  @Effect()
  updateSelfSuccess$ = this.actions$.pipe(
    ofType(UserActionTypes.UPDATE_SELF_SUCCESS),
    map(() => new CreateNotification({
      type: Type.info,
      message: 'Successfully updated your details'
    })));

  @Effect()
  deleteUser$ = combineLatest([
    this.actions$.pipe(ofType<DeleteUser>(UserActionTypes.DELETE)),
    this.store$.select(iamMajorVersion).pipe(filter(identity))])
    .pipe(
      mergeMap(([action, version]: [DeleteUser, IAMMajorVersion]) =>
      this.requests.deleteUser(action.payload, version).pipe(
        map((user: User) => new DeleteUserSuccess(user)),
        catchError((error: HttpErrorResponse) => of(new DeleteUserFailure(error))))));

  @Effect()
  deleteUserSuccess$ = this.actions$.pipe(
    ofType(UserActionTypes.DELETE_SUCCESS),
    map(() => new CreateNotification({
      type: Type.info,
      message: 'Successfully deleted user'
    })));

  @Effect()
  deleteUserFailure$ = this.actions$.pipe(
    ofType(UserActionTypes.DELETE_FAILURE),
    map(({ payload: { error } }: DeleteUserFailure) => {
      const msg = error.error;
      return new CreateNotification({
        type: Type.error,
        message: `Could not delete user: ${msg || error}`
      });
    }));

  @Effect()
  createUser$ = combineLatest([
    this.actions$.pipe(ofType<CreateUser>(UserActionTypes.CREATE)),
    this.store$.select(iamMajorVersion).pipe(filter(identity))])
    .pipe(
      mergeMap(([action, version]: [CreateUser, IAMMajorVersion]) =>
      this.requests.createUser(action.payload, version).pipe(
        map((resp: User) => new CreateUserSuccess(resp)),
        catchError((error) => of(new CreateUserFailure(error))))));

  @Effect()
  createUserSuccess$ = this.actions$.pipe(
    ofType(UserActionTypes.CREATE_SUCCESS),
    map(( { payload: user }: CreateUserSuccess)  =>
      new CreateNotification({
        type: Type.info,
        message: `Created user: ${user.id}.`
      })
    ));

  @Effect()
  createUserFailure$ = this.actions$.pipe(
    ofType<CreateUserFailure>(UserActionTypes.CREATE_FAILURE),
    // username conflict and bad password handled in the modal, see user-management.component.ts
    filter(({ payload: { status } }) => {
      return status !== HttpStatus.CONFLICT && status !== HttpStatus.BAD_REQUEST;
    }),
    map(({ payload: { error } }) => new CreateNotification({
      type: Type.error,
      message: `Could not create user: ${error.error || error}.`
    })));
}
