const path = require("path");
const WorkboxPlugin = require('workbox-webpack-plugin');

module.exports = {
	entry: './src/index.jsx',
	output: {
		path: path.resolve(__dirname, 'public'),
		filename: 'bundle.js',
		publicPath: "/",
	},
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
	plugins: [
		new WorkboxPlugin.GenerateSW({
			// these options encourage the ServiceWorkers to get in there fast
			// and not allow any straggling "old" SWs to hang around
			clientsClaim: true,
			skipWaiting: true
		}),
	],
	resolve: {
		extensions: ['.js', '.jsx'],
	}
};