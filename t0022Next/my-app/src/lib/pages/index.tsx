import { useForm } from "react-hook-form";
import { useRouter } from "next/router";
import Link from 'next/link';
import type { SingData } from '../types/singdata';


export default function EditForm() {
    var defaultValues:SingData = {
        unique_id:0,
        movie:"",
        url:"",
        singStart:"",
        song:"",
    }
        
    return (
        <div>{defaultValues.unique_id}</div>
    );
}
