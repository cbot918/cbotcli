describe 'cbotcli' do
  before do
    puts "[*] setup"
    puts ""
  end

  after do
    puts "[*] teardown"
  end

  def run_script()
    raw_output = nil
    IO.popen("./cbotcli.exe vim","r+") do |pipe|
      raw_output = pipe.gets(nil)
    end
    raw_output.split("\n")
  end

  it 'multicheck GET /m1/script/vim / check POST /m1/signup' do
    result = run_script()
    expect(result).to match_array(["vim.sh",'{"email":"yale@gmail.com", "password":"12345"}','""'])
  end

  

end
