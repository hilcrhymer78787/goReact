// Import the functions you need from the SDKs you need
// import firebase from 'firebase'
import firebase from 'firebase/compat/app';
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
    apiKey: "AIzaSyDpPuLJzwlNKyiMxhIjtMqKGxEALv9EbRE",
    authDomain: "tsumiage-a4563.firebaseapp.com",
    projectId: "tsumiage-a4563",
    storageBucket: "tsumiage-a4563.appspot.com",
    messagingSenderId: "327915970208",
    appId: "1:327915970208:web:b717f373a59f19391e73c5",
    measurementId: "G-DGVCH3PFQX"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);

export default firebase