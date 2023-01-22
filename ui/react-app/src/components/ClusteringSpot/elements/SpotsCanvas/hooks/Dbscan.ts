import { gql } from "@apollo/client";

const QueryDbscan = gql`
  query Dbscan($param: DbscanParam!) {
    dbscan(param: $param) {
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
  }
`;
