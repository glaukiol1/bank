
console.log("Hello From JavaScript File!")

document.getElementById('submit').addEventListener('click',(ev) => {
    const docidnr = document.getElementById('docidnr')
    const password = document.getElementById('password')

    fetch('http://localhost:8080/getuserdata?documentid='+docidnr.value+"&password="+password.value).then(resp=>{
        if (resp.status == 200) {
            resp.json().then(json=>{
                alert(`Hello, ${json.name}`)
            })
        } else {
            resp.json().then(json=>{
                alert(`Error; ${json.message}`)
            })
        }
    })
})