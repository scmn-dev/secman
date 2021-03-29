const decoder = new TextDecoder("utf-8");
const data = await Deno.readFile("map.json");

console.log(decoder.decode(data));
