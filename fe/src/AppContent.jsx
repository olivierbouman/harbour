
import styled from "styled-components";

const Divv = styled.div`
    margin: 40px;
`


const AppContent = ({counter, setCounter}) => {

    const fetchList = () => {
        fetch("https://jsonplaceholder.typicode.com/posts")
        .then(response => response.json())
        .then(json => {
            let posts = document.getElementById("post-list");
            json.forEach(e => {
                let li = document.createElement("li")
                li.appendChild(document.createTextNode(e.title))
                posts.appendChild(li)
            });
        })
    }

    const muisIn = () => {
        setCounter(counter+1)
    }

    return (
        <Divv>
            <button onClick={fetchList} className="btn btn-outline-success">Fetch Data</button>

            <hr />

            <div id="post-list">LIST</div>

            <hr />
            <p onMouseEnter={muisIn}>Moeder muis iss: {counter}</p>

        </Divv>

    );
}


export default AppContent
