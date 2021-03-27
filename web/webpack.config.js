const path = require("path");
const webpack = require("webpack");
const fs = require("fs");

module.exports = (env, options) => {
	const plugins = [
		new webpack.DefinePlugin({
			'process.env': {
				'apiBaseUrl': options.mode === 'development' ? '"http://localhost:8035"' : '"https://spamchamp-ws.gempir.com"',
			}
		}),
	];

	if (options.mode === "production") {
		plugins.push(new HashApplier());
	}

	return {
		entry: './src/index.jsx',
		output: {
			path: path.resolve(__dirname, 'public/bundle'),
			filename: options.mode === "production" ? 'bundle.[hash].js' : "bundle.js",
			publicPath: "/",
		},
		devtool: 'eval-source-map',
		module: {
			rules: [
				{
					test: /\.jsx$/,
					exclude: /node_modules/,
					use: {
						loader: "babel-loader"
					}
				},
				{
					test: /\.s?css$/,
					use: ["style-loader", "css-loader", "sass-loader"]
				},
			]
		},
		stats: {
			// Config for minimal console.log mess.
			assets: false,
			colors: true,
			version: false,
			hash: false,
			modules: false,
			timings: false,
			entrypoints: false,
			chunks: false,
			chunkModules: false
		},
		plugins: plugins,
		resolve: {
			extensions: ['.js', '.jsx'],
		}
	}
}

class HashApplier {
	apply(compiler) {
		compiler.hooks.done.tap('hash-applier', data => {
			const fileContents = fs.readFileSync(__dirname + "/public/index.html", "utf8");
			const newFileContents = fileContents.replace(
				/<!-- webpack-bundle-start -->(.*)?<!-- webpack-bundle-end -->/,
				`<!-- webpack-bundle-start --><script src="/bundle/bundle.${data.hash}.js"></script><!-- webpack-bundle-end -->`
			);

			fs.writeFileSync(__dirname + "/public/index.html", newFileContents);
		});
	}
}