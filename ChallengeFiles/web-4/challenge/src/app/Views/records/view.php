<?php if (! empty($records) && is_array($records)) : ?>
    <div>
        <table class="table is-striped is-hoverable is-fullwidth">
            <tr>
                <th>Type</th>
                <th>From</th>
                <th>To</th>
                <th>Month</th>
                <th>Description</th>
                <th>Amount</th>
                <th>Owner</th>
            </tr>
        <?php foreach ($records as $record): ?>
            <tr>
                <?php if($record['type'] == 0) : ?>
                    <td>Credit</td>
                <?php else : ?>
                    <td>Debit</td>
                <?php endif ?>
                <td><?= $record['who_from'] ?></td>
                <td><?= $record['who_to'] ?></td>
                <td><?= $record['t_month'] ?></td>
                <td><?= $record['description'] ?></td>
                <td>$<?= $record['amount'] ?></td>
                <td><?= $record['owner'] ?></td>
            </tr>
        <?php endforeach; ?>    
        </table>
    </div>
    <!-- Currently disabled whilst we sort out the issues with this functionality. -->
    <div class="columns" style="display:none;">
    <div class="column is-one-third">
    </div>
    <div class="column is-one-third has-text-centered">    
        <form action="/records/view" method="get">
            <label for="filter">Filter by owner: </label>
            <input type="text" name="filter" id="filter" class="input">
            <input type="submit" class="button">
        </form>
    </div>
    <div class="column">
    </div>
    </div>
<?php else : ?>
    <div class="columns">
    <div class="column is-one-third">
    </div>
    <div class="column is-one-third has-text-centered">    
        <h3 class="title is-3">No records</h3>
        <p>Unable to find any records for you.</p>
        <p>Try <a href="/records/create">creating</a> one to get started!</p>
    </div>
    <div class="column">
    </div>
    </div>
<?php endif ?>

