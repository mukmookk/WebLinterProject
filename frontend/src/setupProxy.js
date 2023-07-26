const {
    createProxyMiddleware
} = require('http-proxy-middleware');

module.exports = function (app) {
    app.use(
        '/',
        createProxyMiddleware({
            target: process.env.FORMATTER_SERVICE_PORT,
            changeOrigin: true,

        })
    );
};