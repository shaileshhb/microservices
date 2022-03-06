import { useState } from 'react'
import axios from "axios";
import DisplayComment from './DisplayComment';

function CreateComment({postID}) {

  const [commentMessage, updateCommentMessage] = useState("");

  const onSubmitHandler = async (e) => {
    e.preventDefault();
    console.log(" onSubmitHandler -> ", e);
    console.log(" _postID -> ", postID);
    await axios
      .post(`http://localhost:4002/api/v1/post/${postID}/comments`, {
        title: commentMessage,
      })
      .catch((err) => console.error(err));

    updateCommentMessage("");
  };

  const onSubmitClick = () => {
    if (commentMessage === "" || commentMessage === null) {
      // alert("Post title must be specified")
      return;
    }
    console.log("onSubmitClick -> ", commentMessage);
    updateCommentMessage("");
  };


  return(
    <div>
      <form onSubmit={onSubmitHandler}>
        <div className='form-group'>
          <label>Enter Comment</label>
          <input type="text" className="form-control" placeholder="Post-1" 
            value={commentMessage} onChange={(e) => updateCommentMessage(e.target.value)}/>
        </div>
        <br />
        <button type="submit" className="btn btn-primary" onClick={onSubmitClick}>
          Submit
        </button>
      </form>
      <br />
      <br />
      <DisplayComment/>
    </div>
  )
}

export default CreateComment;