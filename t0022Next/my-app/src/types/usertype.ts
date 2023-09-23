// import { User } from './types'

export type User= {
    MemberId    :   string|null;
    MemberName	:	string;
    Email		:	string|null;
    Password	:	string;
    CreatedAt	:	Date |null;		
  };

export type LoginUser= {
    Email		:	string|null;
    Password	:	string;	
  };