module.exports = {
  configureWebpack: {
    devServer: {
      port: 8081,
      proxy: {
        '/api': {
          target: 'http://localhost:8080',
        },
      },
    },
  },
};
