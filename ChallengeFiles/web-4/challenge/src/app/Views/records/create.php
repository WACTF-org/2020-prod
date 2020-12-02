<div class="columns">
    <div class="column is-one-third">
    </div>
    <div class="column is-one-third">
        <form action="/records/create" class="has-text-centered" method="post">
            <?= csrf_field() ?>

            <?php if(!empty($errors) && is_string($errors)): ?>
                <div class="notification is-danger is-light">
                    <h1 class="title"><?= esc($errors) ?></h1>
                </div>
            <?php elseif(!empty($errors) && is_array($errors)): ?>
                <div class="notification is-danger is-light">
                    <?php foreach ($errors as $error): ?>
                        <h1 class="title"><?= esc($error) ?></h1>
                    <?php endforeach; ?>
                </div>
            <?php endif ?>

            <br>
            <label for="description">Description</label>
            <input type="text" name="description" class="input">
            <br>
            <label for="amount">Amount</label>
            <input type="number" name="amount" id="amount" class="input" min="0">
            <br>
            <label for="who_from">Who From</label>
            <input type="text" name="who_from" id="who_from" class="input">
            <br>
            <label for="who_to">Who To</label>
            <input type="text" name="who_to" id="who_to" class="input">
            <br>
            <label for="type">Type</label>
            <div class="select">
                <select name="type" id="type">
                    <option value="debit">Debit</option>
                    <option value="credit">Credit</option>
                </select>
            </div>
            <br>
            <label for="month">Month</label>
            <div class="select is-multiple">
                <select name="month" id="month" class="select">
                    <option value="1">January</option>
                    <option value="2">February</option>
                    <option value="3">March</option>
                    <option value="4">April</option>
                    <option value="5">May</option>
                    <option value="6">June</option>
                    <option value="7">July</option>
                    <option value="8">August</option>
                    <option value="9">September</option>
                    <option value="10">October</option>
                    <option value="11">November</option>
                    <option value="12">December</option>
                </select>
            </div>
            <br>

            <input type="submit" value="Submit Record" name="submit" class="button">
        </form>
</div>
    <div class="column">
    </div>
</div>