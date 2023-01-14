import { gql, useQuery } from "@apollo/client";
import { SpotsProfile } from "../../../../../generates/types";

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

interface SpotsProfiles {
  spotsProfiles: Array<SpotsProfile>;
}

export const useGetAll = () => {
  const { loading, error, data } = useQuery<SpotsProfiles>(QueryGetAllSpotsProfile);
  const spotsProfiles = data?.spotsProfiles;
  return { loading, error, spotsProfiles };
};

export default useGetAll;
