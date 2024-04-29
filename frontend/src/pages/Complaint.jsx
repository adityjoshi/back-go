
import React, { useState, useEffect } from "react";
import { GetAuthHeader } from "../utils/Headers";
import { Link } from "react-router-dom"; // Import Link from react-router-dom
import clsx from "clsx";

const formatTimestamp = (timestamp) => {
  if (!timestamp) return ''; // Check if timestamp is null or undefined

  const date = new Date(timestamp); 
  if (isNaN(date.getTime())) return ''; // Check if date is invalid

  const options = {
    year: "numeric",
    month: "short",
    day: "numeric",
    hour: "numeric",
    minute: "numeric",
    second: "numeric",
  };
  return new Intl.DateTimeFormat("en-US", options).format(date);
};

const formatTimestamp1 = (timestamp) => {
  if (!timestamp) return ''; // Check if timestamp is null or undefined

  const date = new Date(timestamp); 
  if (isNaN(date.getTime())) return ''; // Check if date is invalid

  const options = {
    year: "numeric",
    month: "short",
    day: "numeric",
  };
  return new Intl.DateTimeFormat("en-US", options).format(date);
};

const ComplaintForm = () => {
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [room, setRoom] = useState("");
  const [complaintIssues, setComplaintType] = useState("electricity");

  const onSubmitForm = async (e) => {
    e.preventDefault();

    if (!name || name.trim() === "") {
      alert("Please enter a valid name.");
      return;
    }
    if (!room || room.trim() === "") {
      alert("Please enter Room No.");
      return;
    }
    if (!description || description.trim() === "") {
      alert("Please enter a valid complaint.");
      return;
    }

    try {
      const headers = GetAuthHeader();
      const body = { name, description, room, complaintIssues }; // Include complaintType in the request body
      const response = await fetch("http://localhost:2426/complaints", {
        method: "POST",
        headers: headers,
        body: JSON.stringify(body),
      });
      window.location = "/";
    } catch (err) {
      console.error(err.message);
    }
  };

  return (
    <>
      <section className="bg-gray-100 py-12 text-gray-800 sm:py-24 h-full">
        <div className="bg-gray-100 mx-auto flex max-w-md flex-col rounded-lg lg:max-w-screen-xl lg:flex-row">
          <div className="max-w-2xl px-4 lg:pr-24">
            <p className="mb-2 text-blue-600">Hostel Grievance Redressal</p>
            <h3 className="mb-5 text-3xl font-semibold">Submit Your Grievance</h3>
            <p className="mb-16 text-md text-gray-600">Hostel Grievance Redressal ensures a swift and confidential resolution of student concerns. We guarantee a quick response to submitted complaints, fostering a secure and comfortable living environment for all hostel residents.</p>
            {/* Existing JSX code... */}
          </div>
          <div className="border border-gray-100 shadow-gray-500/20 mt-8 mb-8 max-w-md bg-white shadow-sm sm:rounded-lg sm:shadow-lg lg:mt-0">
            <div className="relative border-b border-gray-300 p-4 py-8 sm:px-8">
              <h3 className="mb-1 inline-block text-3xl font-medium"><span className="mr-4">Submit Complaint</span><span className="inline-block rounded-md bg-blue-100 px-2 py-1 text-sm text-blue-700 sm:inline">Quick Response</span></h3>
              <p className="text-gray-600">Contact us for hostel grievance redressal</p>
            </div>
            <div className="p-4 sm:p-8">
              <input id="name" type="text" className="mt-1 w-full resize-y overflow-auto rounded-lg border border-gray-300 px-4 py-2 shadow-sm focus:border-blue-500 focus:outline-none hover:border-blue-500" placeholder="Enter Complaint name"  onChange={(e) => setName(e.target.value)}/>
              <input id="email" type="text" className="peer mt-8 w-full resize-y overflow-auto rounded-lg border border-gray-300 px-4 py-2 shadow-sm focus:border-blue-500 focus:outline-none hover:border-blue-500" placeholder="Enter your Room No." onChange={(e) => setRoom(e.target.value)} />
              <label htmlFor="complaintType" className="mt-5 mb-2 inline-block max-w-full">
                Select Complaint Type
              </label>
              <select
                id="complaintType"
                className="mb-8 w-full rounded-lg border border-gray-300 px-4 py-2 shadow-sm focus:border-blue-500 focus:outline-none hover:border-blue-500"
                value={complaintIssues}
                onChange={(e) => setComplaintType(e.target.value)}
              >
                <option value="ELECTRICITY">Electricity</option>
                <option value="WIFI">Wi-Fi</option>
                <option value="other">Other</option>
              </select>
              <label className="mt-5 mb-2 inline-block max-w-full">Tell us about your grievance</label>
              <textarea id="about" className="mb-8 w-full resize-y overflow-auto rounded-lg border border-gray-300 px-4 py-2 shadow-sm focus:border-blue-500 focus:outline-none hover:border-blue-500" onChange={(e) => setDescription(e.target.value)} ></textarea>
              <Link to="/complaints"> {/* Wrap Submit button with Link */}
                <button className="w-full rounded-lg border border-blue-700 bg-blue-700 p-3 text-center font-medium text-white outline-none transition focus:ring hover:border-blue-700 hover:bg-blue-600 hover:text-white" onClick={onSubmitForm}>Submit</button>
              </Link>
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

// const ComplaintsPage = () => {
//   const [complaints, setComplaints] = useState([]);

//   const getComplaints = async (e) => {
//     try {
//       const response = await fetch("http://localhost:2426/complaints", {
//         method: "GET",
//         headers: GetAuthHeader(),
//       });
//       const jsonData = await response.json();

//       setComplaints(jsonData);
//     } catch (err) {
//       console.error(err.message);
//     }
//   };

//   const handleApproval = async (complaint_id) => {
//     try {
//       const response = await fetch(
//         `http://localhost:2426/complaints/:id`,
//         {
//           method: "GET",
//           headers: GetAuthHeader(),
//         }
//       );

//     } catch (err) {
//       console.error(err.message);
//     }
//   };

//   useEffect(() => {
//     getComplaints();
//   }, []);

//   console.log(complaints);

//   return (
//     <div className="bg-gray-100 p-4 sm:p-8 md:p-10 h-screen">
//       <h1 className="text-2xl font-bold mt-20 mb-8">Complaints</h1>
    
//       {complaints === null || complaints.length === 0 ? (
//         <p className="ml-4 mt-2 text-gray-600 text-xl">
//           No complaints registered yet.
//         </p>
//       ) : (
//         <div className="container mx-auto grid gap-8 md:grid-cols-3 sm:grid-cols-1">
//           {complaints.map((complaint) => (
//             <div key={complaint.complaint_id} className="relative flex h-full flex-col rounded-md border border-gray-200 bg-white p-2.5 hover:border-gray-400 sm:rounded-lg sm:p-5">
//               <div className="text-lg mb-2 font-semibold text-gray-900 hover:text-black sm:mb-1.5 sm:text-2xl">
//                 {complaint.name}
//               </div>
//               <p className="text-sm">Created on {formatTimestamp1(complaint.created_at)}</p>
//               <p className="mb-2 mt-2 text-sm font-semibold">Type: {complaint.complaint_issues}</p>
//               <p className="mb-4 text-sm">
//                 {complaint.assigned_at ? `Completed on ${formatTimestamp(complaint.assigned_at)}` : null}
//               </p>
//               <div className="text-md leading-normal text-gray-400 sm:block overflow-hidden" style={{ maxHeight: '100px' }}>
//                 {complaint.description}
//               </div>
              
//               <button
//                 className={clsx(
//                   "group flex w-1/3 mt-3 cursor-pointer items-center justify-center rounded-md px-4 py-2 text-white transition text-sm",
//                   complaint.is_completed ? "bg-green-500" : "bg-red-600"
//                 )}
                
      
//                 onClick={() => handleApproval(complaint.id)}
//               >
//                 <span className="group flex w-full items-center justify-center rounded py-1 text-center font-bold">
//                   {complaint.is_completed ? "Completed" : "Not Completed"}
//                 </span>
//               </button>
//             </div>
//           ))}
//         </div>
//       )}
//       <ComplaintForm />
//     </div>
//   );
// };



const ComplaintsPage = () => {
  const [complaints, setComplaints] = useState([]);

  const getComplaints = async () => {
    try {
      const response = await fetch("http://localhost:2426/complaints", {
        method: "GET",
        headers: GetAuthHeader(),
      });
      const jsonData = await response.json();

      setComplaints(jsonData);
    } catch (err) {
      console.error(err.message);
    }
  };

  useEffect(() => {
    getComplaints();
  }, []);

  // Group complaints by type
  const groupedComplaints = complaints ? complaints.reduce((acc, complaint) => {
    const { complaint_issues } = complaint;
    if (!acc[complaint_issues]) {
      acc[complaint_issues] = [];
    }
    acc[complaint_issues].push(complaint);
    return acc;
  }, {}) : {};

  console.log(groupedComplaints);

  return (
    <div className="bg-gray-100 p-4 sm:p-8 md:p-10 h-screen">
      <h1 className="text-2xl font-bold mt-20 mb-8">Complaints</h1>

      {complaints === null || complaints.length === 0 ? (
        <p className="ml-4 mt-2 text-gray-600 text-xl">No complaints registered yet.</p>
      ) : (
        <div className="container mx-auto grid gap-8 md:grid-cols-3 sm:grid-cols-1">
          {/* Render each type of complaint in its own column */}
          {Object.entries(groupedComplaints).map(([type, complaintsOfType]) => (
            <div key={type} className="flex flex-col">
              <h2 className="text-xl font-bold mb-4">{type}</h2>
              {complaintsOfType.map((complaint) => (
                <div
                  key={complaint.complaint_id}
                  className="relative flex flex-col rounded-md border border-gray-200 bg-white p-2.5 hover:border-gray-400 sm:rounded-lg sm:p-5 mb-4"
                >
                  <div className="text-lg mb-2 font-semibold text-gray-900 hover:text-black sm:mb-1.5 sm:text-2xl">{complaint.name}</div>
                  <p className="text-sm">Created on {formatTimestamp1(complaint.created_at)}</p>
                  <p className="mb-2 mt-2 text-sm font-semibold">Type: {complaint.complaint_issues}</p>
                  <p className="mb-4 text-sm">{complaint.assigned_at ? `Completed on ${formatTimestamp(complaint.assigned_at)}` : null}</p>
                  <div className="text-md leading-normal text-gray-400 sm:block overflow-hidden" style={{ maxHeight: '100px' }}>{complaint.description}</div>
                  <button
                    className={clsx(
                      "group flex w-1/3 mt-3 cursor-pointer items-center justify-center rounded-md px-4 py-2 text-white transition text-sm",
                      complaint.is_completed ? "bg-green-500" : "bg-red-600"
                    )}
                    onClick={() => handleApproval(complaint.id)}
                  >
                    <span className="group flex w-full items-center justify-center rounded py-1 text-center font-bold">
                      {complaint.is_completed ? "Completed" : "Not Completed"}
                    </span>
                  </button>
                </div>
              ))}
            </div>
          ))}
        </div>
      )}
      <ComplaintForm />
    </div>
  );
};

 export default ComplaintsPage;