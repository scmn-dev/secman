import { Octokit } from "octokit";
import { GH_TOKEN } from "../../constants";

const octokit = new Octokit({
  auth: GH_TOKEN,
});

export const GetLatestGHRelease = async (repo: string) => {
  const data = await octokit.rest.repos
    .listReleases({
      owner: "scmn-dev",
      repo: repo,
    })
    .then((res) => {
      return res.data[0].tag_name;
    });

  return data;
};
