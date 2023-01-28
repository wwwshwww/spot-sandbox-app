import { gql, useQuery } from "@apollo/client";
import { Spot } from "../../../../../generated/types";

export const QueryGetAllSpot = gql`
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

export const useGetAllSpot = () => {
  const { loading, error, data } = useQuery<Spots>(QueryGetAllSpot);
  const spots = data?.spots;
  return {loading, error, spots}
}