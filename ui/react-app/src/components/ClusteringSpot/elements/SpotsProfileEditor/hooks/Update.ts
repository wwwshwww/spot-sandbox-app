import { gql } from "@apollo/client";

export const MutationUpdateSpotsProfile = gql`
  mutation UpdateCurrentSpotsProfile($key: Int!, $input: SpotsProfileParam!) {
    updateSpotsProfile(key: $key, input: $input){
      key
      spots{
        key
      }
    }
  }
`;
