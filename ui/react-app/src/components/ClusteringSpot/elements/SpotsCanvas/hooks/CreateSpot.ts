import { gql } from "@apollo/client";

export const MutationCreateSpot = gql`
  mutation MutationCreateSpot($input: LatLng!) {
    createSpot(input: $input){
      key
    }
  }
`;