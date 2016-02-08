kd                  = require 'kd.js'
utils               = require './../../core/utils'
LoginViewInlineForm = require './../../login/loginviewinlineform'
LoginInputView      = require './../../login/logininputview'


module.exports = class TeamLoginAndCreateTabForm extends LoginViewInlineForm

  constructor: (options = {}, data)->

    options.cssClass = 'clearfix login-form'

    super options, data

    { username, email } = utils.getTeamData().signup

    @username = new LoginInputView
      inputOptions        :
        placeholder       : 'Email or username'
        name              : 'username'
        defaultValue      : email or username
        validate          :
          rules           :
            required      : yes
          messages        :
            required      : 'Please enter a username.'

    @password = new LoginInputView
      inputOptions        :
        type              : 'password'
        name              : 'password'
        placeholder       : 'your password'
        validate          :
          rules           :
            required      : yes
          messages        :
            required      : 'Please enter a password.'

    @tfcode = new LoginInputView
      cssClass            : 'hidden'
      inputOptions        :
        name              : 'tfcode'
        placeholder       : 'authentication code'
        testPath          : 'login-form-tfcode'
        attributes        :
          testpath        : 'login-form-tfcode'

    @backLink = new kd.CustomHTMLView
      tagName  : 'span'
      cssClass : 'TeamsModal-button-link back'
      partial  : '<i></i> <a href="/Team/Payment">Back</a>'

    @button = new kd.ButtonView
      title      : 'Sign in'
      style      : 'TeamsModal-button TeamsModal-button--green'
      attributes : testpath : 'register-button'
      type       : 'submit'
      loader     : yes

    @on 'FailedToCreateATeam', @button.bound 'hideLoader'


  showTwoFactor: ->

    @tfcode.show()
    @tfcode.setFocus()


  pistachio: ->

    """
    {{> @username}}{{> @password}}{{> @tfcode}}
    <div class='TeamsModal-button-separator'></div>
    {{> @button}}
    {{> @backLink}}
    """