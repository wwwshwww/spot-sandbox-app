import { gql } from "@apollo/client";
import { ClusterElement } from "../../../../../generated/types";

export const QueryDbscan = gql`
  query Dbscan($param: DbscanParam!) {
    dbscan(param: $param) {
      ClusterElements{
        key
        assignedNumber
        spot {
          key
        }
        paths {
          spot {
            key
          }
        }
      }
      ClusterNum
    }
  }
`;