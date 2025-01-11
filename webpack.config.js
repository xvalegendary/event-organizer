const path = require('path')

module.exports = {
	
	resolve: {
		fallback: {
			http: require.resolve('stream-http'),
			https: require.resolve('https-browserify'),
			stream: require.resolve('stream-browserify'),
			url: require.resolve('url/'),
			util: require.resolve('util/'),
			zlib: require.resolve('browserify-zlib'),
			assert: require.resolve('assert/'),
		},
	},
	
}
