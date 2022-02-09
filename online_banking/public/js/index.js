
console.log("Hello From JavaScript File!")

function clearScreen() {
    document.getElementById('main').innerHTML = ''
}
function updateScreen(data) {
    document.getElementById('main').innerHTML = data
}

function loggedInScreen(data) {
    return `
    <h2>Hello ${data.name}</h2>
    <h3>Available Funds: <br>EUR€${data.addresses[0].balance} ${data.addresses[0].address} <br> USD$${data.addresses[1].balance} ${data.addresses[1].address}</h3>
    `
}
document.getElementById('submit').addEventListener('click',(ev) => {
    const docidnr = document.getElementById('docidnr')
    const password = document.getElementById('password')

    fetch('http://localhost:8080/getuserdata?documentid='+docidnr.value+"&password="+password.value).then(resp=>{
        if (resp.status == 200) {
            resp.json().then(json=>{
                clearScreen()
                updateScreen(loggedInScreen(json))
            })
        } else {
            resp.json().then(json=>{
                alert(`Error; ${json.message}`)
            })
        }
    })
})