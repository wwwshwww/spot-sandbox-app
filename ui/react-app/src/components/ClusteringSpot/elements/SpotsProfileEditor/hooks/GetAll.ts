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

interface Spots {
  spotsProfiles: Array<SpotsProfile>;
}

const useGetAll = () => {
  const { loading, error, data } = useQuery<Spots>(QueryGetAllSpotsProfile);
  const spotsProfiles = data?.spotsProfiles;
  return { loading, error, spotsProfiles };
};

export default useGetAll;
