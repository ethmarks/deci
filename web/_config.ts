import lume from "lume/mod.ts";
import theme from "theme/mod.ts";

const site = lume({
	location: new URL("https://ethmarks.github.io/deci/"),
});

site.use(theme());

export default site;
