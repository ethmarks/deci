// ---
// title: Gallery
// description: screenshots of deci across versions
// layout: layouts/base.vto
// ---

export const title = "Gallery";
export const description = "screenshots of deci across versions";
export const layout = "layouts/base.vto";

const getImg = (file: string): string =>
	`
<h3>${file}</h3>
<p>
<img
alt=${file}
src="/uploads/${file}.png"
image-size
loading="lazy"
decoding="async"
fetchpriority="auto"
class="responsive"/>
</p>
`;

const screenshots: string[] = ["code", "preview", "help", "readme", "empty"];
const screenshotHTML = `
	<section>
	<h2>Screenshots</h2>
	<p>These screenshots are of the current Deci version and were captured using <a href="https://github.com/charmbracelet/vhs">VHS</a> from <a href="https://github.com/ethmarks/deci/blob/main/screenshots.tape"><code>screenshots.tape</code></a></p>
	${screenshots.map(getImg).join("")}
	</section>
	`;

const versions: string[] = [
	"v1.0.0",
	"v0.4.1",
	"v0.3.0",
	"v0.2.2",
	"v0.2.0",
	"v0.1.0",
];
const versionHTML = `
	<section>
	<h2>Versions</h2>
	<p>These screenshots are of past versions of Deci, captured by manually screenshotting my terminal window.</p>
	<p>They don't include any versions before v0.1.0, nor does any of the smaller versions that didn't really have any visual changes.</p>
	<p>For a complete list of versions, see <a href="https://github.com/ethmarks/deci/tags">https://github.com/ethmarks/deci/tags</a>.</p>

	${versions.map(getImg).join("")}
	</section>
	`;

export const content: string = `
<h1>Gallery</h1>

${screenshotHTML}
${versionHTML}
`;
