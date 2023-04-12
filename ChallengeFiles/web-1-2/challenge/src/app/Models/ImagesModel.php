<?php namespace App\Models;

use CodeIgniter\Model;

class ImagesModel extends Model
{
    protected $table = 'files';
    protected $primaryKey = 'id';

    protected $allowedFields = ['description', 'filename', 'owner', 'public'];

    public function createImage($description, $filename, $owner, $public){
        $this->insert(['description' => $description,
                       'filename' => $filename,
                       'owner' => $owner,
                       'public' => $public]);
    }

    public function getImage($id){
        $result = $this->find($id);
        if(sizeof($result) === 0){
            $result = NULL;
        }

        return $result;
    }

    public function getPublicAndOwned($username){
        $images = $this->where('public =', 1)->orWhere('owner =', $username)->findAll();
        return $images;
    }
}