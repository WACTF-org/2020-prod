<?php namespace App\Controllers;

use App\Models\RecordsModel;
use CodeIgniter\Controller;
use mikehaertl\wkhtmlto\Pdf;

class Records extends BaseController
{
    public function index(){
        # Show all the public images.
        if($this->request->getCookie('securelyLoggedInUser') == 'superadmin'){
            # Show the gallery here
            $model = new ImagesModel();
            $data['title'] = "Records";
            $data['images'] = $model->getPublicAndOwned($this->session->username);
            echo view('templates/header', $data);
            echo view('templates/nav');
            echo view('images/gallery', $data);
            echo view('templates/footer', $data);
        }else{
            # They need an account to continue
            return redirect()->to(base_url('users/signup_required'));
        }
    }

    public function create(){
        if($this->request->getCookie('securelyLoggedInUser') == 'superadmin'){
            $data['title'] = 'Create Record';
            $model = new RecordsModel();
            echo view('templates/header', $data);
            echo view('templates/nav');
            if($this->request->getMethod() === 'post'){
                $valid = $this->validate([
                    'description' => 'required|min_length[1]',
                    'amount' => 'required|integer|greater_than_equal_to[0]',
                    'who_from' => 'required|min_length[1]',
                    'who_to' => 'required|min_length[1]',
                    'type' => 'required',
                    'month' => 'required|integer|greater_than_equal_to[1]|less_than_equal_to[12]'
                ]);
                if($valid){
                    if($this->request->getPost('type') === "debit"){
                        $type = TRUE;
                        //debit
                    }else{
                        $type = FALSE;
                        //credit
                    }
                    $model->createRecord(
                        $this->request->getPost('description'),
                        $this->request->getPost('amount'),
                        $this->request->getPost('who_from'),
                        $this->request->getPost('who_to'),
                        $type,
                        $this->request->getPost('month'),
                        $this->request->getCookie('securelyLoggedInUser')
                    );
                    return redirect()->to(base_url('records/view'));
                }else{
                    echo view('records/create', ['errors' => $this->validator->getErrors()]);
                }
            }else{
                echo view('records/create');
            }
            echo view('templates/footer');
        }else{
            # They need an account to continue
            return redirect()->to(base_url('users/signup_required'));
        }
    }

    public function view(){
        if($this->request->getCookie('securelyLoggedInUser') == 'superadmin'){
            $data['title'] = 'View Records';
            $model = new RecordsModel();
            try{
                if($this->request->getGet('filter') !== null){
                    $data['records'] = $model->getRecordsFiltered($this->request->getGet('filter'));
                }else{
                    $data['records'] =  $model->getRecords($this->request->getCookie('securelyLoggedInUser'));
    
                }
                echo view('templates/header', $data);
                echo view('templates/nav');
                echo view('records/view', $data);
                echo view('templates/footer');
            }catch(\Exception $e){
                echo view('templates/header', $data);
                echo view('templates/nav');
                echo '<h1 class="title">' . $e->getMessage() . '</h1>';
                echo view('templates/footer');
            }
        }else{
            # They need an account to continue
            return redirect()->to(base_url('users/signup_required'));
        }
    }
}