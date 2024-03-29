export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type ClusterElement = {
  __typename?: 'ClusterElement';
  assignedNumber: Scalars['Int'];
  dbscanProfile: DbscanProfile;
  key: Scalars['Int'];
  paths: Array<ClusterElement>;
  spot: Spot;
};

/**  TODO: profileのIDを指定するかたちではなく各パラメータを生で扱えるようにすべき  */
export type DbscanParam = {
  dbscanProfileKey: Scalars['Int'];
  spotKeys: Array<Scalars['Int']>;
};

export type DbscanProfile = {
  __typename?: 'DbscanProfile';
  distanceType: DistanceType;
  key: Scalars['Int'];
  maxCount?: Maybe<Scalars['Int']>;
  meterThreshold?: Maybe<Scalars['Int']>;
  minCount: Scalars['Int'];
  minutesThreshold?: Maybe<Scalars['Int']>;
};

export type DbscanProfileParam = {
  distanceType: DistanceType;
  maxCount?: InputMaybe<Scalars['Int']>;
  meterThreshold?: InputMaybe<Scalars['Int']>;
  minCount: Scalars['Int'];
  minutesThreshold?: InputMaybe<Scalars['Int']>;
};

export type DbscanResult = {
  __typename?: 'DbscanResult';
  ClusterElements: Array<ClusterElement>;
  ClusterNum: Scalars['Int'];
};

export enum DistanceType {
  Hubeny = 'Hubeny',
  RouteLength = 'RouteLength',
  TravelTime = 'TravelTime'
}

export type LatLng = {
  lat: Scalars['Float'];
  lng: Scalars['Float'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createDbscanProfile: DbscanProfile;
  createSpot: Spot;
  createSpotsProfile: SpotsProfile;
  updateDbscanProfile: DbscanProfile;
  updateSpotsProfile: SpotsProfile;
};


export type MutationCreateDbscanProfileArgs = {
  input: DbscanProfileParam;
};


export type MutationCreateSpotArgs = {
  input: LatLng;
};


export type MutationCreateSpotsProfileArgs = {
  input: SpotsProfileParam;
};


export type MutationUpdateDbscanProfileArgs = {
  input: DbscanProfileParam;
  key: Scalars['Int'];
};


export type MutationUpdateSpotsProfileArgs = {
  input: SpotsProfileParam;
  key: Scalars['Int'];
};

export type Query = {
  __typename?: 'Query';
  dbscan: DbscanResult;
  dbscanProfile?: Maybe<DbscanProfile>;
  dbscanProfiles: Array<DbscanProfile>;
  spot?: Maybe<Spot>;
  spots: Array<Spot>;
  spotsProfile?: Maybe<SpotsProfile>;
  spotsProfiles: Array<SpotsProfile>;
};


export type QueryDbscanArgs = {
  param: DbscanParam;
};


export type QueryDbscanProfileArgs = {
  key: Scalars['Int'];
};


export type QuerySpotArgs = {
  key: Scalars['Int'];
};


export type QuerySpotsProfileArgs = {
  key: Scalars['Int'];
};

export type Spot = {
  __typename?: 'Spot';
  addressRepr: Scalars['String'];
  key: Scalars['Int'];
  lat: Scalars['Float'];
  lng: Scalars['Float'];
  postalCode: Scalars['String'];
};

export type SpotsProfile = {
  __typename?: 'SpotsProfile';
  key: Scalars['Int'];
  spots: Array<Spot>;
};

export type SpotsProfileParam = {
  spotKeys: Array<Scalars['Int']>;
};
