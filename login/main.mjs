import { initializeApp } from 'firebase/app';
import { getAuth, signInWithCustomToken } from "firebase/auth";

const firebaseConfig = {
  };


initializeApp(firebaseConfig);

const token = "";
const auth = getAuth();
signInWithCustomToken(auth,token)
    .then(res => {
        console.log(res.user.uid);
        res.user.getIdToken(false).then(getIdToken =>{
            console.log(getIdToken); 
        })
        
        // console.log(res.user.refreshToken);
    })
    .catch(e => {
        console.log(e);
    })
