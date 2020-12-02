<?php namespace App\Controllers;

use App\Models\ImagesModel;
use CodeIgniter\Controller;

class Images extends BaseController
{
    public function index(){
        # Show all the public images.
        if(isset($this->session->is_logged_in) && $this->session->is_logged_in){
            # Show the gallery here
            $model = new ImagesModel();
            $data['title'] = "Gallery";
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

    public function upload(){
        # Let a user upload an image.
        if(isset($this->session->is_logged_in) && $this->session->is_logged_in){
            $data['title'] = 'Upload';
            echo view('templates/header', $data);
            echo view('templates/nav');
            if($this->request->getMethod() === 'post'){
                $file = $this->request->getFile('image');
                $valid = $this->validate(['image' => [
                    'uploaded[image]', 
                    'mime_in[image,image/jpg,image/jpeg,image/png]',
                    'is_image[image]',
                    ],
                    'description' => 'required|min_length[1]'
                ]);
                if($valid){
                    $newName = $file->getRandomName();
                    $file->move('/var/www/html/public/uploads', $newName);
                    $model = new ImagesModel();
                    if($this->request->getPost('public') === "on"){
                        $public = TRUE;
                    }else{
                        $public = FALSE;
                    }
                    $model->createImage($this->request->getPost('description'),
                                        $newName,
                                        $this->session->username,
                                        $public);
                    return redirect()->to(base_url('images/'));
                }else{
                    echo view('images/upload', ['errors' => $this->validator->getErrors()]);
                }
            }else{
                echo view('images/upload');
            }
            echo view('templates/footer');
        }else{
            # They need an account to continue
            return redirect()->to(base_url('users/signup_required'));
        }
    }

    public function single($id){
        # Show a single image
        if(isset($this->session->is_logged_in) && $this->session->is_logged_in){
            $model = new ImagesModel();
            if(isset($id)){
                $data['title'] = 'Image - ' . $id;
                echo view('templates/header', $data);
                echo view('templates/nav');
                # they used a key so get their image.
                $data['image'] = $model->getImage($id);
                echo view('images/single', $data);
                echo view('templates/footer');
            }else{
                # they haven't used the id so fuck off and set one please.
                return redirect()->to(base_url('images/'));
            }
        }else{
            # They need an account to continue
            return redirect()->to(base_url('users/signup_required'));
        }
    }
}