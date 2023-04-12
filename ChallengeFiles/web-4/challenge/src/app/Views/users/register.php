<div class="columns">
    <div class="column is-one-third">
    </div>
    <div class="column is-one-third">
        <form action="/users/register" class="has-text-centered" method="post">
            <?= csrf_field() ?>

            <?php if(!empty($errors) && is_string($errors)): ?>
                <div class="notification is-danger is-light">
                    <h1 class="title"><?= esc($errors) ?></h1>
                </div>
            <?php elseif(!empty($errors) && is_array($errors)): ?>
                <div class="notification is-danger is-light">
                    <?php foreach ($errors as $error): ?>
                        <h2 class="title"><?= esc($error) ?></h1>
                    <?php endforeach; ?>
                </div>
            <?php endif ?>

            <label for="username">Username</label>
            <input type="input" name="username" class="input is-rounded"/><br />

            <label for="password">Password</label>
            <input type="password" name="password" class="input is-rounded"></input><br />
            <br>
            <input type="submit" name="submit" value="Sign Up" class="button"/>
            
            <br>
            <h4 class="title is-4">Already have an account? <a href="/users/login">Sign in</a></h4>
        </form>
    </div>
    <div class="column">
    </div>
</div>