import { gql } from "@apollo/client";
import { DbscanProfile } from "../../../../../generated/types";

export const QueryGetAllDbscanProfiles = gql`
  query GetAllDbscanProfile {
    dbscanProfiles {
      key
      distanceType
      minCount
      maxCount
      meterThreshold
      minutesThreshold
    }
  }
`;

export interface DbscanProfiles {
  dbscanProfiles: Array<DbscanProfile>;
}
