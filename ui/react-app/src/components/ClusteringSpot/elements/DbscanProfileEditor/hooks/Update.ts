import { gql } from "@apollo/client";


export const MutationUpdateDbscanProfile = gql`
mutation UpdateDbscanProfile($key: Int!, $input: DbscanProfileParam!) {
  updateDbscanProfile(key: $key, input: $input){
    key
  }
}
`;