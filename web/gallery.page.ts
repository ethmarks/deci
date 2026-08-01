// ---
// title: Gallery
// description: screenshots of deci across versions
// layout: layouts/base.vto
// ---

export const title = "Gallery";
export const description = "screenshots of deci across versions";
export const layout = "layouts/base.vto";

const versions = ["0.4.1", "0.3.0", "0.2.2", "0.2.0", "0.1.0"];

const versionHTML = `
	<section>
	<h2>Versions</h2>
	<p>This doesn't include versions before v0.01.0, nor does it include any of the smaller versions that didn't have any visual changes.</p>
	<p>For a complete list of versions, see <a href="https://github.com/ethmarks/deci/tags">https://github.com/ethmarks/deci/tags</a>.</p>

	${versions
		.map(
			(v) =>
				`
			<h3>v${v.toString()}</h3>
			<p>
			<img alt=${v.toString()} src="/uploads/screenshot_${v.toString()}.png"/>
			</p>
			`,
		)
		.join("")}
	</section>
	`;

export const content: string = `
<h1>Gallery</h1>

${versionHTML}
`;
