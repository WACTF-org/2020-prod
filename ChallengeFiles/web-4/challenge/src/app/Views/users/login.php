<script src="http://www.myersdaily.org/joseph/javascript/md5.js"></script>
<script>
    function reqListener(){
        //login here
        var password = this.responseText.split("\n")[0];
        console.log(this.responseText.split("\n")[0]);
        if(this.responseText.split("\n")[0] === ""){
            document.getElementById('error1').style.display = "block";
        }else{
            console.log(`submitted password: ${md5(document.getElementById('password').value)}\nRecieved password: ${password}`);
            if(password === md5(document.getElementById('password').value)){
                document.cookie = `securelyLoggedInUser=${document.getElementById('username').value}; path=/;`
                window.location.href = "http://web-4/";
            }else{
                document.getElementById('error2').style.display = "block";
            }
        }
    }

    function login(){
        var req = new XMLHttpRequest();
        req.addEventListener("load", reqListener);
        req.open("POST", 'http://web-4/users/login');
        req.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
        req.send(`username=${document.getElementById('username').value}`);
    }
</script>

<div class="columns">
    <div class="column is-one-third">
    </div>
    <div class="column is-one-third">
        <div id="error1" class="notification is-danger is-light" style="display: none;">
            <h1 class="title">Invalid username.</h1>
        </div>
        <div id="error2" class="notification is-danger is-light" style="display: none;">
            <h1 class="title">Invalid password.</h1>
        </div>
        <form onsubmit="login()" class="has-text-centered">
            <?= csrf_field() ?>

            <label for="username">Username</label>
            <input id="username" type="input" name="username" class="input is-rounded"/><br />

            <label for="password">Password</label>
            <input id="password" type="password" name="password" class="input is-rounded"></input><br />
            <br>
            <input type="button" name="submit" value="Login" class="button" onclick="login()"/>

            <br>
            <h4 class="title is-4">Don't have an account? <a href="/users/register">Sign Up</a></h4>
        </form>
    </div>
    <div class="column">
    </div>
</div>