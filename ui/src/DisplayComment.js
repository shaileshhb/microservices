
function DisplayComment({comments}) {
  // const [comments, updateComments] = useState({});

  // const loadComments = async () => {
  //   const response = await axios
  //     .get(`http://localhost:4002/api/v1/post/${postID}/comments`)
  //     .catch((err) => console.error(err));

  //   console.log(response);
  //   updateComments(response.data);
  // };

  // useEffect(() => {
  //   loadComments();
  //   updateComments(comments)
  // }, []); // [] -> callback only once

  console.log(comments);

  const cardOfComments = Object.values(comments).map((c) => {
    return (
      <div className="d-flex justify-content-between" key={c.id}>
        <div className="card">
          <div className="card-body">
            <h3>{c.message}</h3>
          </div>
        </div>
      </div>
    );
  });

  return <div className="container">{cardOfComments}</div>;
}

export default DisplayComment;
