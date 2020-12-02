<?php namespace App\Controllers;

class Home extends BaseController
{
	public function index()
	{
		if($this->request->getCookie('securelyLoggedInUser') == 'superadmin'){
			$data['title'] = 'Home';
			echo view('templates/header', $data);
			echo view('templates/nav');
			echo view('welcome');
			echo view('templates/footer');
        }else{
           # They need an account to continue
           return redirect()->to(base_url('users/signup_required'));
        }
	}

	//--------------------------------------------------------------------

}
