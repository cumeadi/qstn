
module.exports = {
  entry: './src/index',
  output: {
    filename: 'bundle.js',
    path: 'static',
  },
  module: {
    loaders: [{
      exclude: /node_modules/,
      loader: 'babel',
      test: /\.js$/,
    }],
  }
};

