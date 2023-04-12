window.onload = function() {
    var apiKeyInput = document.getElementById('api-key-input')
    var sigInput = document.getElementById('sig-input')
    var codeInput = document.getElementById('code-input')
    var signBtn = document.getElementById('sign-btn')
    var runBtn = document.getElementById('run-btn')
    var signError = document.getElementById('sig-error')
    var resultText = document.getElementById('result-text')

    signBtn.addEventListener('click', function() {
        if (apiKeyInput.value === '') {
            signError.innerText = 'No API key specified'
        } else {
            signError.innerText = ''
            sigInput.value = md5(apiKeyInput.value + '|' + codeInput.value)
        }
    })

    runBtn.addEventListener('click', function(e) {
        e.preventDefault()
        var sig = sigInput.value
        var code = codeInput.value
        
        var r = new XMLHttpRequest()
        r.onreadystatechange = handleResponse
        r.open('POST', '/')
        r.setRequestHeader('Content-Type', 'application/json')  
        r.send(JSON.stringify({ sig, code: btoa(code) }))

        function handleResponse() {
            try {
                if (r.readyState === XMLHttpRequest.DONE) {
                    if (r.status === 200) {
                        var res = JSON.parse(r.responseText)
                        resultText.value = res.result
                    } else {
                        console.log('error', r)
                    }
                }
            } catch(e) {
                console.error(e)
            }
        }
    })
}
