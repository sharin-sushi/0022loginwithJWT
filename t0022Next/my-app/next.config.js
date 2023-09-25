// https用..ここに書くのは一般的ではない？
// const { readFileSync } = require('fs');
// const { join } = require('path');

// const httpsOptions = {
//   key: readFileSync(join(__dirname, 'key/server.key')),
//   cert: readFileSync(join(__dirname, 'key/server.cert')),
// };

// module.exports = {
//   server: {
//     https: httpsOptions,
//   },
// };

// ******* memo *******


// 通常はこれで良い多分
/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
}

module.exports = nextConfig



// フロントとバックを同じオリジンとみなす…。
// module.exports = {
//   async rewrites() {
//     return [
//       {
//         source: '/api/:path*',
//         destination: 'http://localhost:8080/:path*',
//       },
//     ];
//   },
// };