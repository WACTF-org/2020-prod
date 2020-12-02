<?php namespace App\Models;

use CodeIgniter\Model;

class RecordsModel extends Model
{
    protected $table = 'records';
    protected $primaryKey = 'id';

    protected $allowedFields = ['description', 'amount', 'who_from', 'who_to', 'type', 't_month', 'owner'];

    public function createRecord($description, $amount, $who_from, $who_to, $type, $t_month, $owner){
        $this->insert(['description' => $description, 
                       'amount' => $amount, 
                       'who_from' => $who_from, 
                       'who_to' => $who_to, 
                       'type' => $type, 
                       't_month' => $t_month, 
                       'owner' => $owner]);
    }

    public function getRecords($owner){
        $result = $this->where('owner = ', $owner)->findAll();
        return $result;
    }

    public function getRecordsFiltered($owner){
        $db = \Config\Database::connect();
        $query = $db->query("SELECT * FROM records WHERE owner='" . $owner . "'");
        $results = $query->getResultArray();
        return $results;
    }

    public function deleteRecord($id){
        $this->delete($id);
    }

    public function updateRecord($id, $description, $amount, $who_from, $who_to, $type, $t_month, $owner){
        $this->update($id, ['description' => $description, 
                            'amount' => $amount, 
                            'who_from' => $who_from, 
                            'who_to' => $who_to, 
                            'type' => $type, 
                            't_month' => $t_month, 
                            'owner' => $owner]);
    }
}