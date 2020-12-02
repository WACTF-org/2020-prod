<?php namespace App\Models;

use CodeIgniter\Model;

class UsersModel extends Model
{
    protected $table = 'users';
    protected $primaryKey = 'id';

    protected $allowedFields = ['username', 'password'];

    public function getUsers($user = false){
        if($user === false){
            return $this->findAll();
        }

        return $this->asArray()
                    ->where(['username' => $user])
                    ->first();
    }

    public function getUser($id){
        # I think it is dumb that I need to reference index 0 on something
        # that is only ever going to return 1 row but whatever.
        $result = $this->find($id);
        if(sizeof($result) === 0){
            $result = NULL;
        }

        return $result;
    }

    public function getUserPass($username){
        $user = $this->asArray()->where(['username' => $username])->first();
        return $user['password'];
    }

    public function getUserId($username){
        $user = $this->asArray()->where(['username' => $username])->first();
        return $user['id'];
    }

    public function createUser($username, $password){
        $this->insert(['username' => $username, 'password' => $password]);
    }
}
