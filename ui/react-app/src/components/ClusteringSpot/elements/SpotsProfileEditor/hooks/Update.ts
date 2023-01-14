import { gql } from "@apollo/client";

const QueryGetAllSpotsProfile = gql`
query GetAllSpotsProfile {
  spotsProfiles {
    key
    spots {
      key
    }
  }
}
`;