import isLoggedIn from "../lib/auth"
import { VFC, useEffect, useState } from "react"
import type { User } from "../types/usertype"
import { useRouter } from "next/router";
import { getAllCookies } from "../lib/getallcookie";


type WithGetAccessControl<P> = P & {
  getAccessControl?: GetAccessControl;
}

const MyPage: WithGetAccessControl<VFC<{ data: User }>> = ({ data }) => {
  const router = useRouter();

  useEffect(() => {
    if (!data) {
      router.push('/signin');
    }
  }, [data]);

  return (
    <div>
      <h2>会員情報</h2>
      <ul>
        <a>memberId：</a>{data.MemberId}<br />
        <a>memberName:</a>{data.MemberName} <br />
        <a>Password：</a>{data.Password}<br />
        <a>email：</a>{data.Email}<br />
        <li>createdAt: {data.CreatedAt?.toString()}</li>
        {/* <a>createdAt　　　　：</a>{data.CreatedAt}<br /> */}
      </ul>
    </div>
  );
}

// アクセス制御
MyPage.getAccessControl = () => {
  return isLoggedIn() ? { type: 'replace', destination: '/signin' } : null
}

export default MyPage;

export const getServerSideProps = async (context: { req: { headers: { cookie: any; }; }; }) => {
  const rawCookie = context.req.headers.cookie;
  console.log("rawCookie=", rawCookie, "\n")
  const sessionToken = rawCookie?.split(';').find(cookie => cookie.trim().startsWith('token='))?.split('=')[1];
  console.log("sessionToken=", sessionToken, "\n")
  const options: RequestInit = {
    headers: {
      cookie: `token=${sessionToken}`,
    },
    cache: "no-store",
    credentials: 'include'
  };
  const res = await fetch('http://localhost:8080/users/profile', options);
  const data = await res.json();

  return {
    props: {
      data
    }
  }
}




  
// ****memo****
  {/* <Link href={`${data.url}&t=${data.singStart}`}>動画サイト</><br />
  {`↑${data.url}&t=${data.singStart}`}<br></br>
  <button color="red" >戻る </button>
  <Link href="/">戻る </>
  <Link href={`/edit?Unique_id=${data.unique_id}`}>編集</Link><br /> */}