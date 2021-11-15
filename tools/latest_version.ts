import packageJson from "package-json";

export default async function latestVersion(packageName: any, options: any) {
  const { version } = await packageJson(packageName.toLowerCase(), options);
  return version;
}
