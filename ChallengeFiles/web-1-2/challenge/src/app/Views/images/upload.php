<div class="columns">
    <div class="column is-one-third">
    </div>
    <div class="column is-one-third">
        <form action="upload" class="has-text-centered" method="post" enctype="multipart/form-data">
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

            <!-- <label for="image">Select image to upload:</label> -->
            <!-- <input type="file" class="" name="image" id="image"> -->
            
            <div class="file is-boxed" style="justify-content:center;">
                <label class="file-label">
                    <input class="file-input" type="file" name="image">
                    <span class="file-cta">
                    <span class="file-icon">
                        <i class="fas fa-upload"></i>
                    </span>
                    <span class="file-label">
                        Choose a file…
                    </span>
                    </span>
                </label>
            </div>

            <br>
            <label for="description">Description</label>
            <input type="text" name="description" class="input"><br>
            <label for="public">Publicly visible?</label>
            <input type="checkbox" name="public" id="public" class="checkbox">
            <br>

            <input type="submit" value="Upload Image" name="submit" class="button">
        </form>
</div>
    <div class="column">
    </div>
</div>