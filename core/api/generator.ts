import { COMPLEXIES } from "../constants";

const generator = COMPLEXIES.filter((item) => item.checked).reduce(
  (acc, current) => {
    return acc + current.value;
  },
  ""
);

export default generator;
