import { useEffect, useState } from 'react';
import Link from 'next/link';
// import YouTube from 'react-youtube';
import style from '../Youtube.module.css';
import type { AllData, Streamer, StreamerMovie } from '../types/singdata'; //type{}で型情報のみインポート
// import DeleteButton from '../components/DeleteButton';
import { useRouter } from 'next/router';
import https from 'https';
import axios from 'axios';

//分割代入？
// 型注釈IndexPage(posts: Post)
const AllDatePage = ({ posts }) =>  {
    console.log("test2")

    // data1というステートを定義。streamerの配列を持つ。
    // setData1はステートを更新する関数。
  const [streamers, setData1] = useState<Streamer[]>();
  const [streamerstoMovies, setData2] = useState<StreamerMovie[]>();
  const router = useRouter();

  useEffect(() => {
    if (posts) {
        setData1(posts.streamers);
        setData2(posts.streamerstoMovies);

    }
}, [posts]);
console.log("test3")

  return (
    <div>
      <h1>TOP画面</h1><br />
        <h3>"推し"の"歌枠"の聴きたい"歌"を再生しよう。 <br />
        推しが歌った"歌"を一目で把握、布教しよう。
        </h3><br />

           {/*  ログイン機能のリンクボタン */}
            <Link href="/signup"><button style={{ background: 'darkblue' }}>
             会員登録</button>     &nbsp;
            </Link>
            <Link href="/signin"><button style={{ background: 'darkblue' }}>       
                ログイン</button>  &nbsp;    
            </Link>
            <Link href="/mypage"><button style={{ background: 'brown' }}>
                マイページ</button>
            </Link>
            <br />　　　　　↑ゲストログイン可能です <br /><br /><br />

        {/* 一覧表示 */}

        {/* 配信者一覧 */}
        <h2>★配信者</h2>
        <Link href={`/create`} ><u>推し登録</u></Link>

      <table border={4} >
        <thead> {/* ← tabeleのhead */}
          <tr>
            <td>ID</td>
            <td>推し</td>
            <td>読み</td>
            <td>紹介動画</td>
            <td>リンク</td>
            <td>リンク</td>
            <td>編集</td>
          </tr>
        </thead>
    
        <tbody>
        {streamers && streamers.map((Data1, index) => (
        <tr key={index}>
         <td>{Data1.StreamerId}</td>
         <td>{Data1.StreamerName}</td>
         <td>{Data1.NameKana}</td>
       
              {Data1.SelfIntroUrl ? (
          <td><Link href={Data1.SelfIntroUrl}>youtubeへ</Link></td>
        ) : (
          <td>未登録</td>
        )}

          <td><Link href={`/movie?streamer_id=${Data1.StreamerId}`}>歌枠</Link></td>
          <td><Link href={`/sing?streamer_id=${Data1.StreamerId}`}>歌</Link></td>
        
              {/* http://localhost:3000/show?Unique_id=1　になった */}
              <td><Link href={`/edit?Unique_id=${Data1.StreamerId}`}>編集</Link></td>
            </tr>
            ))}
        </tbody>
      </table><br />

      {/*動画一覧  */}
      <h2>★動画</h2>
      <Link href={`/create`} ><u>歌枠を登録</u></Link>
      <table border={4} >
        <thead>
           <tr>
            <td>配信者名</td>
            <td>動画ID</td>
            <td>動画名(クリックで視聴)</td>
            <td>動画url</td>
            <td>編集</td>
          </tr>
        </thead>
        <tbody>
          {streamerstoMovies && streamerstoMovies.map((item2, index) => (
            <tr key={index}>
              <td>{item2.StreamerName}</td>
              <td>{item2.MovieId}</td>        
              <td><Link href={`/sing?movie_url=${item2.MovieUrl}`}>{item2.MovieTitle}</Link></td>
              <td>{item2.MovieUrl}</td>
 
              {/* http://localhost:3000/show?Unique_id=1　になった */}
              <td><Link href={`/posts/${item2.StreamerId}`}>編集</Link></td>
            </tr>
            ))}
        </tbody>
      </table><br />
    </div>
  )};

//   オレオレ証明書対応のために、証明書を無視してる。
// 本番環境では変更が必要。
  export async function getServerSideProps() {
    const httpsAgent = new https.Agent({
        rejectUnauthorized: false
    });
    let resData;
    try {
        const response = await axios.get('https://localhost:8080', {
            httpsAgent: process.env.NODE_ENV === "production" ? undefined : httpsAgent
        });
        resData = response.data;
    } catch (error) {
        // throw new Error('Network response has err');
        console.log("axios.getでerroe:", error)
    }

    // console.log("resData: ", resData)
    console.log("test1")

    return {
        props: {
            posts: resData
        }
    };
}
  
  export default AllDatePage;
  

//   **** memo ****

//      <DeleteButton Unique_id={item.streamer_id} /> 

// 各種リンク
            //  {item2.Movies ? ( <td>{item2.Movies'MovieId'}</td>
            //   ) : ( <td>未登録</td>      )}

            //  {item2.MovieUrl ? ( <td>{item2.MovieUrl}</td>
            //   ) : (<td>-</td>            )}
            // {item2.MovieTitle ? (  <td>{item2.MovieTitle}</td>
            //   ) : (<td>-</td>            )} 

// SSR化する前のコード
// function AllDatePage( posts : AllData)  {
//     // data1というステートを定義。streamerの配列を持つ。
//     // setData1はステートを更新する関数。
//   const [streamers, setData1] = useState<Streamer[]>();
//   const [streamerstoMovies, setData2] = useState<StreamerMovie[]>();
//   const router = useRouter();

//   useEffect(() => {  
//     async function fetchData() {
//         try {
//             const response = await fetch('https://localhost:8080');
//                 if (!response.ok) {
//                     throw new Error('Network response was not ok');
//                 }
//             const resData = await response.json();
//             setData1(resData.streamers);
//             setData2(resData.streamers_and_moviesmovies);
//         } catch (error) {
//             console.error("There was a problem with the fetch operation:", error);
//         }
//     }
//     fetchData();
// }, []);

//   return (