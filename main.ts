const main = async () => {
  (await import(`./days/${Deno.args[0]}/index.ts`))?.default();
};

main();
