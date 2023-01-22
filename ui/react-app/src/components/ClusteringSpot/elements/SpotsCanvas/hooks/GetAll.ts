import { gql, useQuery } from "@apollo/client";
import { Spot } from "../../../../../generated/types";

const QueryGetAllSpot = gql`
  query GetAllSpot {
    spots {
      key
      lat
      lng
      addressRepr
    }
  }
`;

interface Spots {
  spots: Array<Spot>;
}

export const useGetAll = () => {
  const { loading, error, data } = useQuery<Spots>(QueryGetAllSpot);
  const spots = data?.spots;
  return {loading, error, spots}
}