import { gql } from "@apollo/client";

export const MutationUpdateSpotsProfile = gql`
  mutation UpdateSpotsProfile($key: Int!, $input: SpotsProfileParam!) {
    updateSpotsProfile(key: $key, input: $input){
      key
      spots{
        key
      }
    }
  }
`;
