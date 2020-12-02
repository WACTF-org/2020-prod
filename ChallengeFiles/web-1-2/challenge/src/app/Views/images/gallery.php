<?php if (! empty($images) && is_array($images)) : ?>
    <div>
        <?php foreach ($images as $image): ?>
            <div class="container">
                <a href="/images/view/<?= $image['id']?>">
                    <img class="image" src="/uploads/<?= $image['filename'] ?>" style="max-width: 600px; margin-left: auto; margin-right:auto; border: solid grey 0.2rem; border-radius: 5%;">
                </a>
                <h3 class="subtitle is-3 has-text-centered" style="padding-bottom:0.2rem;"><?= $image['description'] ?></h3>
            </div>
        <?php endforeach; ?>    
    </div>

<?php else : ?>
    <div class="columns">
    <div class="column is-one-third">
    </div>
    <div class="column is-one-third has-text-centered">    
        <h3 class="title is-3">No images</h3>
        <p>Unable to find any images for you.</p>
        <p>Try <a href="images/upload">uploading</a> one to get started!</p>
    </div>
    <div class="column">
    </div>
    </div>

<?php endif ?>