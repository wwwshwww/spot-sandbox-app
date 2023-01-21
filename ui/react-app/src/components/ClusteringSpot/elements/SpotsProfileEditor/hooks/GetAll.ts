import { gql, useApolloClient, useQuery } from "@apollo/client";
import { SpotsProfile } from "../../../../../generated/types";

export const QueryGetAllSpotsProfile = gql`
  query GetAllSpotsProfile {
    spotsProfiles {
      key
      spots {
        key
      }
    }
  }
`;

export interface SpotsProfiles {
  spotsProfiles: Array<SpotsProfile>;
}

export const useGetAll = () => {
  const { loading, error, data } = useQuery<SpotsProfiles>(
    QueryGetAllSpotsProfile
  );
  const spotsProfiles = data?.spotsProfiles;
  return { loading, error, spotsProfiles };
};

export default useGetAll;
