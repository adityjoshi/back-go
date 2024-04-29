
import Navbar from "./Navbar";
import { useState, useEffect } from "react";
import { GetAuthHeader } from "../utils/Headers";

function AccountPage() {
  const [userDetails, setUserDetails] = useState({
    full_name: "",
    email: "",
    phone: "",
    usn: "",
    block_id: "",
    block_name: "",
    room: "",
  });

  const [userType, setUserType] = useState(null);

  useEffect(() => {
    const fetchUserType = async () => {
      try {
        const response = await fetch("http://localhost:2426/userType", {
          method: "GET",
          headers: GetAuthHeader(),
        });

        if (response.ok) {
          const data = await response.json();
          setUserType(data.userType);
        } else {
          console.error('Failed to fetch user type');
        }
      } catch (error) {
        console.error(error.message);
      }
    };

    fetchUserType();
  }, []); 

  useEffect(() => {
    const getUserDetails = async (id) => {
      try {
        const response = await fetch("http://localhost:2426/userDetails/${id}", {
          method: "GET",
          headers: GetAuthHeader(),
        });
        if (response.ok) {
          const data = await response.json();
          console.log("User details:", data); // Log received data
          setUserDetails(data);
        } else {
          console.error('Failed to fetch user details');
        }
      } catch (error) {
        console.error(error.message);
      }
    };
  
    getUserDetails();
  }, []);
  
    


return (
  <>
    <Navbar />
    <h2 className="mt-20 ml-5 mr-5 text-2xl font-semibold">Profile</h2>

    <ul className="mt-6 flex flex-col ml-5 mr-5 ">
      <li className="lg:w-1/3 sm:w-full inline-flex items-center gap-x-2 py-3 px-4 text-sm border text-gray-800 -mt-px first:rounded-t-lg first:mt-0 last:rounded-b-lg">
        <div className="flex items-center justify-between w-full">
          <span>Name</span>
          <span>{userDetails.full_name}</span>
        </div>
      </li>
      <li className="lg:w-1/3 sm:w-full inline-flex items-center gap-x-2 py-3 px-4 text-sm border text-gray-800 -mt-px first:rounded-t-lg first:mt-0 last:rounded-b-lg">
        <div className="flex items-center justify-between w-full">
          <span>Email</span>
          <span>{userDetails.email}</span>
        </div>
      </li>
      <li className="lg:w-1/3 sm:w-full inline-flex items-center gap-x-2 py-3 px-4 text-sm border text-gray-800 -mt-px first:rounded-t-lg first:mt-0 last:rounded-b-lg">
        <div className="flex items-center justify-between w-full">
          <span>Phone</span>
          <span>{userDetails.phone}</span>
        </div>
      </li>
      {userType !== 'warden' && (
        <>
          <li className="lg:w-1/3 sm:w-full inline-flex items-center gap-x-2 py-3 px-4 text-sm border text-gray-800 -mt-px first:rounded-t-lg first:mt-0 last:rounded-b-lg">
            <div className="flex items-center justify-between w-full">
              <span>USN</span>
              <span>{userDetails.usn}</span>
            </div>
          </li>
          <li className="lg:w-1/3 sm:w-full inline-flex items-center gap-x-2 py-3 px-4 text-sm border text-gray-800 -mt-px first:rounded-t-lg first:mt-0 last:rounded-b-lg">
            <div className="flex items-center justify-between w-full">
              <span>Block ID</span>
              <span>{userDetails.block_id}</span>
            </div>
          </li>
          {/* <li className="lg:w-1/3 sm:w-full inline-flex items-center gap-x-2 py-3 px-4 text-sm border text-gray-800 -mt-px first:rounded-t-lg first:mt-0 last:rounded-b-lg">
            <div className="flex items-center justify-between w-full">
              <span>Block Name</span>
              <span>{userDetails.block_name}</span>
            </div>
          </li> */}
          <li className="lg:w-1/3 sm:w-full inline-flex items-center gap-x-2 py-3 px-4 text-sm border text-gray-800 -mt-px first:rounded-t-lg first:mt-0 last:rounded-b-lg">
            <div className="flex items-center justify-between w-full">
              <span>Room</span>
              <span>{userDetails.room}</span>
            </div>
          </li>
        </>
      )}
    </ul> 
    <button className="mt-5 ml-5 relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-cyan-500 to-blue-500 group-hover:from-cyan-500 group-hover:to-blue-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-cyan-200 dark:focus:ring-cyan-800">
      <a className="relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-blue-500 rounded-md group-hover:bg-opacity-0" href="/">
        Back
      </a>
    </button>
  </>
);
}
export default AccountPage;
