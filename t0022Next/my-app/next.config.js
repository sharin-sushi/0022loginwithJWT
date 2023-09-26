// 通常はこれで良い多分
/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
}

module.exports = nextConfig


// ******* memo *******

// フロントとバックを同じオリジンとみなす…。
// 別の方法で解決させた。多分、SameSite None, Secure trueでhttps化
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